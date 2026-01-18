/**
 * Type guard functions for runtime type checking.
 * Use these to narrow `unknown` types safely.
 */

import type { User, Admin, Address, Status, STATUS } from "./types-example";

// =============================================================================
// PRIMITIVE GUARDS
// =============================================================================

export function isString(value: unknown): value is string {
  return typeof value === "string";
}

export function isNumber(value: unknown): value is number {
  return typeof value === "number" && !Number.isNaN(value);
}

export function isBoolean(value: unknown): value is boolean {
  return typeof value === "boolean";
}

export function isObject(value: unknown): value is Record<string, unknown> {
  return typeof value === "object" && value !== null && !Array.isArray(value);
}

export function isArray(value: unknown): value is unknown[] {
  return Array.isArray(value);
}

export function isNonEmptyString(value: unknown): value is string {
  return isString(value) && value.trim().length > 0;
}

export function isNonEmptyArray<T>(value: T[]): value is [T, ...T[]] {
  return value.length > 0;
}

// =============================================================================
// DOMAIN GUARDS
// =============================================================================

/**
 * Check if value is a valid Status
 */
export function isStatus(value: unknown): value is Status {
  return (
    isString(value) &&
    Object.values(STATUS).includes(value as Status)
  );
}

/**
 * Check if value is a valid Address
 */
export function isAddress(value: unknown): value is Address {
  if (!isObject(value)) return false;
  
  return (
    isNonEmptyString(value.street) &&
    isNonEmptyString(value.city) &&
    isNonEmptyString(value.state) &&
    isNonEmptyString(value.zipCode) &&
    isNonEmptyString(value.country)
  );
}

/**
 * Check if value is a valid User
 */
export function isUser(value: unknown): value is User {
  if (!isObject(value)) return false;
  
  return (
    isNonEmptyString(value.id) &&
    isNonEmptyString(value.email) &&
    isNonEmptyString(value.name) &&
    isString(value.role) &&
    isStatus(value.status) &&
    isAddress(value.address)
  );
}

/**
 * Check if value is an Admin (User with admin role)
 */
export function isAdmin(value: unknown): value is Admin {
  if (!isUser(value)) return false;
  
  return (
    value.role === "admin" &&
    "permissions" in value &&
    isArray((value as Admin).permissions)
  );
}

/**
 * Check if value is an array of Users
 */
export function isUserArray(value: unknown): value is User[] {
  return isArray(value) && value.every(isUser);
}

// =============================================================================
// ASSERTION FUNCTIONS
// =============================================================================

/**
 * Assert value is a User, throw if not
 */
export function assertIsUser(value: unknown): asserts value is User {
  if (!isUser(value)) {
    throw new TypeError(
      `Expected User, got ${typeof value}: ${JSON.stringify(value)}`
    );
  }
}

/**
 * Assert value is defined (not null/undefined)
 */
export function assertDefined<T>(
  value: T | null | undefined,
  message = "Value is null or undefined"
): asserts value is T {
  if (value === null || value === undefined) {
    throw new TypeError(message);
  }
}

/**
 * Assert condition is true
 */
export function assert(
  condition: boolean,
  message = "Assertion failed"
): asserts condition {
  if (!condition) {
    throw new Error(message);
  }
}

// =============================================================================
// EXHAUSTIVENESS CHECK
// =============================================================================

/**
 * Use in switch default case to ensure all cases are handled
 */
export function assertNever(value: never, message?: string): never {
  throw new Error(message ?? `Unexpected value: ${JSON.stringify(value)}`);
}

// Usage example:
// switch (status) {
//   case "active": return "Active";
//   case "inactive": return "Inactive";
//   default: return assertNever(status);
// }

// =============================================================================
// SAFE PARSING
// =============================================================================

/**
 * Safely parse JSON with type validation
 */
export function safeParseJSON<T>(
  json: string,
  guard: (value: unknown) => value is T
): T | null {
  try {
    const parsed = JSON.parse(json);
    return guard(parsed) ? parsed : null;
  } catch {
    return null;
  }
}

/**
 * Parse JSON and throw if invalid
 */
export function parseJSONOrThrow<T>(
  json: string,
  guard: (value: unknown) => value is T,
  errorMessage = "Invalid JSON structure"
): T {
  const result = safeParseJSON(json, guard);
  if (result === null) {
    throw new TypeError(errorMessage);
  }
  return result;
}

// Usage:
// const user = safeParseJSON(jsonString, isUser);
// if (user) { /* safely use user */ }
