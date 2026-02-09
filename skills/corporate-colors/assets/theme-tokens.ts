/**
 * Corporate Color Palette - Based on Catppuccin
 * 
 * This file provides a complete theme system with light (Latte) 
 * and dark (Mocha) mode support.
 * 
 * Usage:
 *   import { theme } from './theme-tokens'
 *   const buttonColor = theme.colors.primary.main
 */

export const catppuccinColors = {
  latte: {
    // Base colors
    base: '#eff1f5',
    mantle: '#e6e9ef',
    crust: '#dce0e8',
    
    // Text colors
    text: '#4c4f69',
    subtext1: '#5c5f77',
    subtext0: '#6c6f85',
    
    // Overlay colors
    overlay2: '#9ca0b0',
    overlay1: '#bcc0cc',
    overlay0: '#acb0be',
    
    // Surface colors
    surface2: '#acb0be',
    surface1: '#bcc0cc',
    surface0: '#ccd0da',
    
    // Brand colors
    blue: '#1e66f5',
    lavender: '#7287fd',
    sapphire: '#04a5e5',
    sky: '#209fb5',
    teal: '#179299',
    green: '#40a02b',
    yellow: '#df8e1d',
    peach: '#fe640b',
    maroon: '#e64553',
    red: '#d20f39',
    mauve: '#8839ef',
    pink: '#ea76cb',
  },
  mocha: {
    // Base colors
    base: '#1e1e2e',
    mantle: '#181825',
    crust: '#11111b',
    
    // Text colors
    text: '#cdd6f4',
    subtext1: '#bac2de',
    subtext0: '#a6adc8',
    
    // Overlay colors
    overlay2: '#9399b2',
    overlay1: '#7f849c',
    overlay0: '#6c7086',
    
    // Surface colors
    surface2: '#585b70',
    surface1: '#45475a',
    surface0: '#313244',
    
    // Brand colors
    blue: '#89b4fa',
    lavender: '#b4befe',
    sapphire: '#74c7ec',
    sky: '#89dceb',
    teal: '#94e2d5',
    green: '#a6e3a1',
    yellow: '#f9e2af',
    peach: '#fab387',
    maroon: '#eba0ac',
    red: '#f38ba8',
    mauve: '#cba6f7',
    pink: '#f5c2e7',
  },
} as const

/**
 * Theme configuration with semantic color mappings
 */
export const theme = {
  light: {
    primary: {
      main: catppuccinColors.latte.blue,
      hover: catppuccinColors.latte.sapphire,
      active: catppuccinColors.latte.lavender,
    },
    secondary: {
      main: catppuccinColors.latte.mauve,
      hover: catppuccinColors.latte.pink,
      muted: catppuccinColors.latte.overlay2,
    },
    accent: {
      success: catppuccinColors.latte.green,
      warning: catppuccinColors.latte.yellow,
      error: catppuccinColors.latte.red,
      info: catppuccinColors.latte.teal,
    },
    surface: {
      base: catppuccinColors.latte.base,
      mantle: catppuccinColors.latte.mantle,
      crust: catppuccinColors.latte.crust,
      card: catppuccinColors.latte.mantle,
      elevated: catppuccinColors.latte.crust,
    },
    text: {
      primary: catppuccinColors.latte.text,
      secondary: catppuccinColors.latte.subtext1,
      tertiary: catppuccinColors.latte.subtext0,
      disabled: catppuccinColors.latte.overlay2,
    },
    border: {
      default: catppuccinColors.latte.overlay0,
      subtle: catppuccinColors.latte.overlay1,
      emphasis: catppuccinColors.latte.mauve,
    },
  },
  dark: {
    primary: {
      main: catppuccinColors.mocha.blue,
      hover: catppuccinColors.mocha.sapphire,
      active: catppuccinColors.mocha.lavender,
    },
    secondary: {
      main: catppuccinColors.mocha.mauve,
      hover: catppuccinColors.mocha.pink,
      muted: catppuccinColors.mocha.overlay2,
    },
    accent: {
      success: catppuccinColors.mocha.green,
      warning: catppuccinColors.mocha.yellow,
      error: catppuccinColors.mocha.red,
      info: catppuccinColors.mocha.teal,
    },
    surface: {
      base: catppuccinColors.mocha.base,
      mantle: catppuccinColors.mocha.mantle,
      crust: catppuccinColors.mocha.crust,
      card: catppuccinColors.mocha.mantle,
      elevated: catppuccinColors.mocha.surface0,
    },
    text: {
      primary: catppuccinColors.mocha.text,
      secondary: catppuccinColors.mocha.subtext1,
      tertiary: catppuccinColors.mocha.subtext0,
      disabled: catppuccinColors.mocha.overlay2,
    },
    border: {
      default: catppuccinColors.mocha.overlay0,
      subtle: catppuccinColors.mocha.overlay1,
      emphasis: catppuccinColors.mocha.mauve,
    },
  },
} as const

/**
 * Type-safe theme getter
 */
export type Theme = typeof theme.light
export type ThemeMode = 'light' | 'dark'

export const getTheme = (mode: ThemeMode): Theme => {
  return theme[mode]
}

/**
 * CSS variable generator
 */
export const generateCSSVariables = (mode: ThemeMode): string => {
  const colors = theme[mode]
  
  return `
    /* Primary */
    --color-primary: ${colors.primary.main};
    --color-primary-hover: ${colors.primary.hover};
    --color-primary-active: ${colors.primary.active};
    
    /* Secondary */
    --color-secondary: ${colors.secondary.main};
    --color-secondary-hover: ${colors.secondary.hover};
    --color-secondary-muted: ${colors.secondary.muted};
    
    /* Accent */
    --color-success: ${colors.accent.success};
    --color-warning: ${colors.accent.warning};
    --color-error: ${colors.accent.error};
    --color-info: ${colors.accent.info};
    
    /* Surface */
    --color-bg-base: ${colors.surface.base};
    --color-bg-mantle: ${colors.surface.mantle};
    --color-bg-crust: ${colors.surface.crust};
    --color-bg-card: ${colors.surface.card};
    --color-bg-elevated: ${colors.surface.elevated};
    
    /* Text */
    --color-text-primary: ${colors.text.primary};
    --color-text-secondary: ${colors.text.secondary};
    --color-text-tertiary: ${colors.text.tertiary};
    --color-text-disabled: ${colors.text.disabled};
    
    /* Border */
    --color-border: ${colors.border.default};
    --color-border-subtle: ${colors.border.subtle};
    --color-border-emphasis: ${colors.border.emphasis};
  `.trim()
}
