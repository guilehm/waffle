import { ThemeProvider } from "@mui/material/styles"
import { AppRouterCacheProvider } from "@mui/material-nextjs/v13-appRouter"
import NavBar from "@/components/NavBar/NavBar"
import Footer from "@/components/Footer/Footer"
import { Raleway } from "next/font/google"

import theme from "../theme"
import "./globals.css"


const raleway = Raleway({ subsets: ["latin"] })

export const metadata = {
  title: "Waffle",
  description: "",
}

export default function RootLayout(props) {
  return (
    <html lang="en">
      <body className={raleway.className}>
        <AppRouterCacheProvider>
          <ThemeProvider theme={theme}>
            <NavBar>
            </NavBar>
            {props.children}
            <Footer />
          </ThemeProvider>
        </AppRouterCacheProvider>
      </body>
    </html >
  )
}
