/**
 * Example type definitions following the const types pattern
 * and flat interface conventions.
 */

// =============================================================================
// CONST TYPES PATTERN
// =============================================================================

/**
 * Status values - const object is the single source of truth
 */
export const STATUS = {
  ACTIVE: "active",
  INACTIVE: "inactive",
  PENDING: "pending",
  ARCHIVED: "archived",
} as const;

export type Status = (typeof STATUS)[keyof typeof STATUS];

/**
 * HTTP status codes
 */
export const HTTP_STATUS = {
  OK: 200,
  CREATED: 201,
  NO_CONTENT: 204,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  NOT_FOUND: 404,
  INTERNAL_ERROR: 500,
} as const;

export type HttpStatus = (typeof HTTP_STATUS)[keyof typeof HTTP_STATUS];

/**
 * User roles with labels
 */
export const USER_ROLE = {
  ADMIN: "admin",
  EDITOR: "editor",
  VIEWER: "viewer",
} as const;

export type UserRole = (typeof USER_ROLE)[keyof typeof USER_ROLE];

export const USER_ROLE_LABELS: Record<UserRole, string> = {
  [USER_ROLE.ADMIN]: "Administrator",
  [USER_ROLE.EDITOR]: "Editor",
  [USER_ROLE.VIEWER]: "Viewer",
};

// =============================================================================
// FLAT INTERFACES
// =============================================================================

/**
 * Address - extracted to keep User flat
 */
export interface Address {
  street: string;
  city: string;
  state: string;
  zipCode: string;
  country: string;
}

/**
 * Base user interface - flat, one level deep
 */
export interface User {
  id: string;
  email: string;
  name: string;
  role: UserRole;
  status: Status;
  address: Address;
  createdAt: Date;
  updatedAt: Date;
}

/**
 * Admin extends User with additional permissions
 */
export interface Admin extends User {
  role: typeof USER_ROLE.ADMIN;
  permissions: string[];
  lastLogin: Date;
}

/**
 * API payload types using utility types
 */
export type CreateUserPayload = Omit<User, "id" | "createdAt" | "updatedAt">;
export type UpdateUserPayload = Partial<Omit<User, "id" | "createdAt" | "updatedAt">>;
export type UserPreview = Pick<User, "id" | "name" | "email" | "role">;

// =============================================================================
// API RESPONSE TYPES
// =============================================================================

/**
 * Generic API response wrapper
 */
export interface ApiResponse<T> {
  data: T;
  status: HttpStatus;
  message: string;
  timestamp: string;
}

/**
 * Paginated response with metadata
 */
export interface PaginatedResponse<T> {
  data: T[];
  meta: {
    total: number;
    page: number;
    pageSize: number;
    totalPages: number;
  };
}

/**
 * Error response
 */
export interface ApiError {
  status: HttpStatus;
  message: string;
  code: string;
  details?: Record<string, string[]>;
}

// =============================================================================
// DISCRIMINATED UNIONS
// =============================================================================

/**
 * Async state for data fetching
 */
export interface IdleState {
  status: "idle";
}

export interface LoadingState {
  status: "loading";
}

export interface SuccessState<T> {
  status: "success";
  data: T;
}

export interface ErrorState {
  status: "error";
  error: Error;
}

export type AsyncState<T> = IdleState | LoadingState | SuccessState<T> | ErrorState;

/**
 * Result type for operations that can fail
 */
export type Result<T, E = Error> =
  | { ok: true; value: T }
  | { ok: false; error: E };

// =============================================================================
// GENERIC UTILITIES
// =============================================================================

/**
 * Make specific properties required
 */
export type RequireFields<T, K extends keyof T> = T & Required<Pick<T, K>>;

/**
 * Make specific properties optional
 */
export type OptionalFields<T, K extends keyof T> = Omit<T, K> & Partial<Pick<T, K>>;

/**
 * Deep readonly
 */
export type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};

/**
 * Non-empty array
 */
export type NonEmptyArray<T> = [T, ...T[]];

/**
 * Nullable type helper
 */
export type Nullable<T> = T | null;

/**
 * Dictionary/Record with string keys
 */
export type Dictionary<T> = Record<string, T>;
