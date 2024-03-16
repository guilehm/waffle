"use client"

import Container from '@mui/material/Container'
import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import Link from '@mui/material/Link'
import Box from '@mui/material/Box'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import { useTheme } from '@emotion/react'



export default function Footer() {

  const theme = useTheme()
  console.log("useTheme", theme)

  return (
    <Box
      component="footer"
      sx={{
        backgroundColor: theme.palette.primary.lightest,
        py: 6,
      }}
    >
      <Container maxWidth="lg">
        <Typography variant="h5" align="center" gutterBottom>
          Download our app
        </Typography>
        <Typography variant="subtitle1" align="center" component="p">
          Stay connected
        </Typography>

        <Typography variant="body2" align="center" component="p" sx={{ mt: 4 }}>
          © Beautiful Footer, 2024
        </Typography>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            mt: 2,
          }}
        >
          <Link href="#" variant="body2" sx={{ mx: 1 }}>
            About us
          </Link>
          |
          <Link href="#" variant="body2" sx={{ mx: 1 }}>
            Contact us
          </Link>
          |
          <Link href="#" variant="body2" sx={{ mx: 1 }}>
            Privacy Policy
          </Link>
        </Box>
      </Container>
    </Box>
  )
}
