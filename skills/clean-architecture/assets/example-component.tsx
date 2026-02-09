/**
 * Example Component - Corporate Colors in Practice
 * 
 * This demonstrates how to use the corporate color system
 * in a React component with TypeScript.
 */

import React from 'react'
import { getTheme, type ThemeMode } from './theme-tokens'

interface ButtonProps {
  variant?: 'primary' | 'secondary' | 'success' | 'warning' | 'error'
  children: React.ReactNode
  onClick?: () => void
  disabled?: boolean
}

export const Button: React.FC<ButtonProps> = ({
  variant = 'primary',
  children,
  onClick,
  disabled = false,
}) => {
  // Get theme colors (you would typically get this from context)
  const mode: ThemeMode = 'light' // or from useTheme() hook
  const colors = getTheme(mode)

  // Map variant to color
  const variantColors = {
    primary: {
      bg: colors.primary.main,
      bgHover: colors.primary.hover,
      bgActive: colors.primary.active,
      text: colors.surface.base,
    },
    secondary: {
      bg: colors.secondary.main,
      bgHover: colors.secondary.hover,
      bgActive: colors.secondary.main,
      text: colors.surface.base,
    },
    success: {
      bg: colors.accent.success,
      bgHover: colors.accent.success,
      bgActive: colors.accent.success,
      text: colors.surface.base,
    },
    warning: {
      bg: colors.accent.warning,
      bgHover: colors.accent.warning,
      bgActive: colors.accent.warning,
      text: colors.text.primary,
    },
    error: {
      bg: colors.accent.error,
      bgHover: colors.accent.error,
      bgActive: colors.accent.error,
      text: colors.surface.base,
    },
  }

  const buttonColors = variantColors[variant]

  return (
    <button
      onClick={onClick}
      disabled={disabled}
      style={{
        backgroundColor: disabled ? colors.secondary.muted : buttonColors.bg,
        color: disabled ? colors.text.disabled : buttonColors.text,
        border: `1px solid ${colors.border.subtle}`,
        borderRadius: '6px',
        padding: '8px 16px',
        fontSize: '14px',
        fontWeight: 600,
        cursor: disabled ? 'not-allowed' : 'pointer',
        transition: 'all 0.2s ease',
        opacity: disabled ? 0.6 : 1,
      }}
      onMouseEnter={(e) => {
        if (!disabled) {
          e.currentTarget.style.backgroundColor = buttonColors.bgHover
        }
      }}
      onMouseLeave={(e) => {
        if (!disabled) {
          e.currentTarget.style.backgroundColor = buttonColors.bg
        }
      }}
      onMouseDown={(e) => {
        if (!disabled) {
          e.currentTarget.style.backgroundColor = buttonColors.bgActive
        }
      }}
      onMouseUp={(e) => {
        if (!disabled) {
          e.currentTarget.style.backgroundColor = buttonColors.bgHover
        }
      }}
    >
      {children}
    </button>
  )
}

/**
 * Card component example
 */
interface CardProps {
  title: string
  children: React.ReactNode
  elevated?: boolean
}

export const Card: React.FC<CardProps> = ({
  title,
  children,
  elevated = false,
}) => {
  const mode: ThemeMode = 'light'
  const colors = getTheme(mode)

  return (
    <div
      style={{
        backgroundColor: elevated ? colors.surface.elevated : colors.surface.card,
        border: `1px solid ${colors.border.default}`,
        borderRadius: '8px',
        padding: '16px',
        boxShadow: elevated ? '0 4px 6px rgba(0, 0, 0, 0.1)' : 'none',
      }}
    >
      <h3
        style={{
          color: colors.text.primary,
          fontSize: '18px',
          fontWeight: 600,
          marginBottom: '12px',
        }}
      >
        {title}
      </h3>
      <div
        style={{
          color: colors.text.secondary,
          fontSize: '14px',
        }}
      >
        {children}
      </div>
    </div>
  )
}

/**
 * Alert component example
 */
interface AlertProps {
  variant: 'success' | 'warning' | 'error' | 'info'
  children: React.ReactNode
}

export const Alert: React.FC<AlertProps> = ({ variant, children }) => {
  const mode: ThemeMode = 'light'
  const colors = getTheme(mode)

  const variantConfig = {
    success: {
      bg: colors.accent.success,
      icon: '✓',
    },
    warning: {
      bg: colors.accent.warning,
      icon: '⚠',
    },
    error: {
      bg: colors.accent.error,
      icon: '✕',
    },
    info: {
      bg: colors.accent.info,
      icon: 'ℹ',
    },
  }

  const config = variantConfig[variant]

  return (
    <div
      style={{
        backgroundColor: `${config.bg}15`, // 15 = ~8% opacity in hex
        border: `1px solid ${config.bg}`,
        borderRadius: '6px',
        padding: '12px 16px',
        display: 'flex',
        alignItems: 'center',
        gap: '12px',
      }}
    >
      <span
        style={{
          color: config.bg,
          fontSize: '18px',
          fontWeight: 700,
        }}
      >
        {config.icon}
      </span>
      <div
        style={{
          color: colors.text.primary,
          fontSize: '14px',
        }}
      >
        {children}
      </div>
    </div>
  )
}

/**
 * Usage Example
 */
export const ExampleUsage: React.FC = () => {
  return (
    <div style={{ padding: '24px', display: 'flex', flexDirection: 'column', gap: '16px' }}>
      <h1>Corporate Colors Example</h1>

      <section>
        <h2>Buttons</h2>
        <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
          <Button variant="primary">Primary Action</Button>
          <Button variant="secondary">Secondary Action</Button>
          <Button variant="success">Success</Button>
          <Button variant="warning">Warning</Button>
          <Button variant="error">Error</Button>
          <Button variant="primary" disabled>Disabled</Button>
        </div>
      </section>

      <section>
        <h2>Cards</h2>
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '16px' }}>
          <Card title="Standard Card">
            This is a standard card with default styling.
          </Card>
          <Card title="Elevated Card" elevated>
            This card has elevated styling with a shadow.
          </Card>
        </div>
      </section>

      <section>
        <h2>Alerts</h2>
        <div style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
          <Alert variant="success">Operation completed successfully!</Alert>
          <Alert variant="warning">Please review your input carefully.</Alert>
          <Alert variant="error">An error occurred while processing.</Alert>
          <Alert variant="info">Here's some helpful information.</Alert>
        </div>
      </section>
    </div>
  )
}
