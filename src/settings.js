export const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost.com/api"

export const settings = {
  siteName: "Fuzzy Train",
  pages: [
    {
      title: "Home",
      url: "/",
    },
  ],
  profile: [
    {
      title: "Profile",
      url: "",
    },
    {
      title: "Logout",
      url: "",
    },
  ],
  footer: {
    about: "",
    description: "",
  },
}
