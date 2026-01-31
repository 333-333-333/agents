---
name: typescript
description: >
  TypeScript strict patterns and best practices.
  Trigger: When implementing or refactoring TypeScript in .ts/.tsx (types, interfaces, generics, const maps, type guards, removing any, tightening unknown).
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [root]
  auto_invoke:
    - "Writing TypeScript types or interfaces"
    - "Refactoring to remove any types"
    - "Creating type guards or generics"
---

## Critical Rules

1. **Const types pattern** - ALWAYS const object first, then extract type
2. **Flat interfaces** - Max one level deep, nested → dedicated interface
3. **No `any`** - Use `unknown` + type guards or generics
4. **Import types** - Use `import type` for type-only imports

---

## Const Types Pattern (REQUIRED)

```typescript
// ✅ ALWAYS: Create const object first, then extract type
const STATUS = {
  ACTIVE: "active",
  INACTIVE: "inactive",
  PENDING: "pending",
} as const;

type Status = (typeof STATUS)[keyof typeof STATUS];
// Result: "active" | "inactive" | "pending"

// ✅ Use the const for runtime, type for annotations
function setStatus(status: Status) {
  if (status === STATUS.ACTIVE) { /* ... */ }
}

// ❌ NEVER: Direct union types (no runtime values)
type Status = "active" | "inactive" | "pending";
```

**Why?** Single source of truth, runtime values, autocomplete, easier refactoring.

### Const Pattern with Labels

```typescript
const HTTP_STATUS = {
  OK: 200,
  CREATED: 201,
  BAD_REQUEST: 400,
  NOT_FOUND: 404,
} as const;

type HttpStatus = (typeof HTTP_STATUS)[keyof typeof HTTP_STATUS];
// Result: 200 | 201 | 400 | 404

// With display labels
const STATUS_LABELS: Record<Status, string> = {
  [STATUS.ACTIVE]: "Active",
  [STATUS.INACTIVE]: "Inactive",
  [STATUS.PENDING]: "Pending",
};
```

---

## Flat Interfaces (REQUIRED)

```typescript
// ✅ ALWAYS: One level depth, nested objects → dedicated interface
interface UserAddress {
  street: string;
  city: string;
  zipCode: string;
}

interface User {
  id: string;
  name: string;
  email: string;
  address: UserAddress;  // Reference, not inline
}

// ✅ Extend for variations
interface Admin extends User {
  permissions: string[];
  role: "admin";
}

interface Guest extends Pick<User, "id" | "name"> {
  sessionId: string;
}

// ❌ NEVER: Inline nested objects
interface User {
  address: { street: string; city: string };  // NO!
}

// ❌ NEVER: Deep nesting
interface User {
  profile: {
    settings: {
      notifications: { email: boolean };  // NO!
    };
  };
}
```

---

## Never Use `any`

```typescript
// ✅ Use unknown for truly unknown types
function parseJSON(input: string): unknown {
  return JSON.parse(input);
}

// ✅ Narrow with type guards
function processUser(input: unknown): User {
  if (isUser(input)) return input;
  throw new Error("Invalid user data");
}

// ✅ Use generics for flexible types
function first<T>(arr: T[]): T | undefined {
  return arr[0];
}

// ✅ For event handlers, type the event
function handleClick(event: React.MouseEvent<HTMLButtonElement>) {
  // ...
}

// ❌ NEVER use any
function parse(input: any): any { }  // NO!
const data: any = fetchData();       // NO!
```

### Dealing with Third-Party `any`

```typescript
// ✅ Wrap and type external APIs
interface ExternalApiResponse {
  data: User[];
  meta: { total: number };
}

async function fetchUsers(): Promise<ExternalApiResponse> {
  const response = await externalApi.get("/users");
  return response as ExternalApiResponse;  // Type assertion at boundary
}

// ✅ Use type assertion only at boundaries, validate if critical
function fromExternalApi(data: unknown): User {
  if (!isUser(data)) {
    throw new TypeError("Invalid user from API");
  }
  return data;
}
```

---

## Type Guards

```typescript
// ✅ Basic type guard
function isString(value: unknown): value is string {
  return typeof value === "string";
}

// ✅ Object type guard
function isUser(value: unknown): value is User {
  return (
    typeof value === "object" &&
    value !== null &&
    "id" in value &&
    "name" in value &&
    typeof (value as User).id === "string" &&
    typeof (value as User).name === "string"
  );
}

// ✅ Array type guard
function isUserArray(value: unknown): value is User[] {
  return Array.isArray(value) && value.every(isUser);
}

// ✅ Discriminated union guard
interface Success { status: "success"; data: User }
interface Error { status: "error"; message: string }
type Result = Success | Error;

function isSuccess(result: Result): result is Success {
  return result.status === "success";
}
```

### Assertion Functions

```typescript
// ✅ Assert and throw if invalid
function assertIsUser(value: unknown): asserts value is User {
  if (!isUser(value)) {
    throw new TypeError("Expected User");
  }
}

// Usage - type is narrowed after assertion
function process(input: unknown) {
  assertIsUser(input);
  console.log(input.name);  // TypeScript knows it's User
}
```

---

## Utility Types

### Selection & Exclusion

```typescript
// Pick specific fields
type UserPreview = Pick<User, "id" | "name">;

// Omit specific fields  
type UserWithoutId = Omit<User, "id">;

// Combine for API payloads
type CreateUserPayload = Omit<User, "id" | "createdAt">;
type UpdateUserPayload = Partial<Omit<User, "id">>;
```

### Modifiers

```typescript
// All fields optional
type PartialUser = Partial<User>;

// All fields required
type RequiredUser = Required<User>;

// All fields readonly
type ReadonlyUser = Readonly<User>;

// Deep readonly (custom)
type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};
```

### Object Types

```typescript
// Record for dictionaries
type UserById = Record<string, User>;
type StatusCount = Record<Status, number>;

// Index signatures
interface Cache {
  [key: string]: User | undefined;  // Always include undefined
}
```

### Union Manipulation

```typescript
type Status = "active" | "inactive" | "pending" | "deleted";

// Extract subset
type ActiveStatus = Extract<Status, "active" | "pending">;
// Result: "active" | "pending"

// Exclude subset
type VisibleStatus = Exclude<Status, "deleted">;
// Result: "active" | "inactive" | "pending"

// Remove null/undefined
type NonNullUser = NonNullable<User | null | undefined>;
```

### Function Types

```typescript
function createUser(name: string, email: string): User {
  return { id: crypto.randomUUID(), name, email };
}

// Extract return type
type CreateUserReturn = ReturnType<typeof createUser>;
// Result: User

// Extract parameters
type CreateUserParams = Parameters<typeof createUser>;
// Result: [string, string]

// First parameter
type FirstParam = Parameters<typeof createUser>[0];
// Result: string
```

---

## Generics

### Basic Generics

```typescript
// ✅ Generic function
function identity<T>(value: T): T {
  return value;
}

// ✅ Generic with constraint
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}

// ✅ Generic interface
interface ApiResponse<T> {
  data: T;
  status: number;
  message: string;
}

// Usage
type UserResponse = ApiResponse<User>;
type UsersResponse = ApiResponse<User[]>;
```

### Generic Constraints

```typescript
// Must have id property
interface HasId {
  id: string;
}

function findById<T extends HasId>(items: T[], id: string): T | undefined {
  return items.find(item => item.id === id);
}

// Must be object
function merge<T extends object, U extends object>(a: T, b: U): T & U {
  return { ...a, ...b };
}
```

### Default Generic Types

```typescript
interface PaginatedResponse<T, M = { total: number }> {
  data: T[];
  meta: M;
}

// Uses default meta
type UserPage = PaginatedResponse<User>;

// Custom meta
type UserPageWithCursor = PaginatedResponse<User, { cursor: string }>;
```

---

## Discriminated Unions

```typescript
// ✅ Use literal type as discriminant
interface LoadingState {
  status: "loading";
}

interface SuccessState<T> {
  status: "success";
  data: T;
}

interface ErrorState {
  status: "error";
  error: Error;
}

type AsyncState<T> = LoadingState | SuccessState<T> | ErrorState;

// ✅ Exhaustive handling
function handleState<T>(state: AsyncState<T>): string {
  switch (state.status) {
    case "loading":
      return "Loading...";
    case "success":
      return `Got ${state.data}`;
    case "error":
      return `Error: ${state.error.message}`;
    default:
      // Exhaustiveness check
      const _exhaustive: never = state;
      return _exhaustive;
  }
}
```

---

## Import Types

```typescript
// ✅ Type-only imports (removed at compile time)
import type { User, Status } from "./types";

// ✅ Mixed imports
import { createUser, type Config } from "./utils";

// ✅ Re-export types
export type { User, Status } from "./types";
export { createUser } from "./utils";
```

---

## File Organization

```
src/
├── types/
│   ├── index.ts          # Re-exports all types
│   ├── user.ts           # User-related types
│   ├── api.ts            # API response types
│   └── constants.ts      # Const maps and derived types
├── utils/
│   └── guards.ts         # Type guard functions
└── features/
    └── users/
        └── types.ts      # Feature-specific types
```

### types/index.ts

```typescript
// Re-export all public types
export type { User, Admin, Guest } from "./user";
export type { ApiResponse, PaginatedResponse } from "./api";
export { STATUS, HTTP_STATUS } from "./constants";
export type { Status, HttpStatus } from "./constants";
```

---

## tsconfig.json Recommendations

```json
{
  "compilerOptions": {
    "strict": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "noImplicitOverride": true,
    "forceConsistentCasingInFileNames": true,
    "verbatimModuleSyntax": true
  }
}
```

| Option | Why |
|--------|-----|
| `strict` | Enables all strict checks |
| `noUncheckedIndexedAccess` | `arr[0]` returns `T \| undefined` |
| `exactOptionalPropertyTypes` | `{ x?: string }` doesn't accept `undefined` |
| `verbatimModuleSyntax` | Enforces `import type` syntax |

---

## Quick Reference

| Pattern | Use |
|---------|-----|
| `const X = {} as const` | Enum-like values with types |
| `(typeof X)[keyof typeof X]` | Extract union from const |
| `interface A extends B` | Inheritance |
| `Pick<T, K>` | Select fields |
| `Omit<T, K>` | Exclude fields |
| `Partial<T>` | All optional |
| `Record<K, V>` | Object/dictionary type |
| `value is Type` | Type guard return |
| `asserts value is Type` | Assertion function |
| `T extends Constraint` | Generic constraint |

## Resources

- **Templates**: See [assets/](assets/) for tsconfig and type examples
