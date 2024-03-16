import { ThemeProvider } from "@mui/material/styles"
import { AppRouterCacheProvider } from "@mui/material-nextjs/v13-appRouter"

import theme from "../theme"

import "./globals.css"
import NavBar from "@/components/NavBar/NavBar"

export const metadata = {
  title: "Fuzzy Train",
  description: "",
}

export default function RootLayout(props) {
  return (
    <html lang="en">
      <body>
        <AppRouterCacheProvider>
          <ThemeProvider theme={theme}>
            <NavBar>
              {props.children}
            </NavBar>
          </ThemeProvider>
        </AppRouterCacheProvider>
      </body>
    </html>
  )
}
