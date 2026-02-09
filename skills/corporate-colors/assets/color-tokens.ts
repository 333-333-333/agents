/**
 * Corporate Color Tokens — Catppuccin Warm Tones
 *
 * Semantic color mappings for light (Latte) and dark (Mocha) modes.
 * Based on warm Catppuccin colors: Maroon, Flamingo, Peach, Rosewater.
 *
 * Usage:
 *   import { primary, secondary, accent, surface, text, border } from './color-tokens'
 */

// =============================================================================
// PRIMARY (Maroon)
// =============================================================================

export const primary = {
  light: {
    main: '#E64553',      // Latte Maroon — primary actions
    hover: '#FE640B',     // Latte Peach — hover state (emphasis)
    active: '#DD7878',    // Latte Flamingo — active/pressed
  },
  dark: {
    main: '#EBA0AC',      // Mocha Maroon — primary actions
    hover: '#FAB387',     // Mocha Peach — hover state (emphasis)
    active: '#F2CDCD',    // Mocha Flamingo — active/pressed
  }
} as const

// =============================================================================
// SECONDARY (Flamingo)
// =============================================================================

export const secondary = {
  light: {
    main: '#DD7878',      // Latte Flamingo — secondary actions
    hover: '#DC8A78',     // Latte Rosewater — hover state
    muted: '#9ca0b0',     // Latte Overlay2 — disabled state
  },
  dark: {
    main: '#F2CDCD',      // Mocha Flamingo — secondary actions
    hover: '#F5E0DC',     // Mocha Rosewater — hover state
    muted: '#9399b2',     // Mocha Overlay2 — disabled state
  }
} as const

// =============================================================================
// ACCENT / SEMANTIC
// =============================================================================

export const accent = {
  light: {
    success: '#40a02b',   // Latte Green — success states
    warning: '#df8e1d',   // Latte Yellow — warnings
    error: '#d20f39',     // Latte Red — errors
    info: '#DC8A78',      // Latte Rosewater — informational (warm)
  },
  dark: {
    success: '#a6e3a1',   // Mocha Green — success states
    warning: '#f9e2af',   // Mocha Yellow — warnings
    error: '#f38ba8',     // Mocha Red — errors
    info: '#F5E0DC',      // Mocha Rosewater — informational (warm)
  }
} as const

// =============================================================================
// SURFACE
// =============================================================================

export const surface = {
  light: {
    base: '#eff1f5',      // Latte Base — main background
    mantle: '#e6e9ef',    // Latte Mantle — secondary background
    crust: '#dce0e8',     // Latte Crust — elevated surfaces
    surface0: '#ccd0da',  // Latte Surface0 — cards
    surface1: '#bcc0cc',  // Latte Surface1 — hovered cards
    surface2: '#acb0be',  // Latte Surface2 — active surfaces
  },
  dark: {
    base: '#1e1e2e',      // Mocha Base — main background
    mantle: '#181825',    // Mocha Mantle — secondary background
    crust: '#11111b',     // Mocha Crust — elevated surfaces
    surface0: '#313244',  // Mocha Surface0 — cards
    surface1: '#45475a',  // Mocha Surface1 — hovered cards
    surface2: '#585b70',  // Mocha Surface2 — active surfaces
  }
} as const

// =============================================================================
// TEXT
// =============================================================================

export const text = {
  light: {
    primary: '#4c4f69',   // Latte Text — primary text
    secondary: '#5c5f77', // Latte Subtext1 — secondary text
    tertiary: '#6c6f85',  // Latte Subtext0 — tertiary/muted
    carbon: '#2B2B2B',    // Near-black — headlines, strong emphasis
  },
  dark: {
    primary: '#cdd6f4',   // Mocha Text — primary text
    secondary: '#bac2de', // Mocha Subtext1 — secondary text
    tertiary: '#a6adc8',  // Mocha Subtext0 — tertiary/muted
    carbon: '#E8E8E8',    // Soft white — headlines, strong emphasis
  }
} as const

// =============================================================================
// BORDER
// =============================================================================

export const border = {
  light: {
    default: '#acb0be',   // Latte Overlay0 — default borders
    subtle: '#bcc0cc',    // Latte Overlay1 — subtle borders
    emphasis: '#E64553',  // Latte Maroon — emphasized borders
  },
  dark: {
    default: '#6c7086',   // Mocha Overlay0 — default borders
    subtle: '#7f849c',    // Mocha Overlay1 — subtle borders
    emphasis: '#EBA0AC',  // Mocha Maroon — emphasized borders
  }
} as const
