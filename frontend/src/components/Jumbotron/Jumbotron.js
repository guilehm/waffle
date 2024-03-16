import { Box, Typography, Button } from "@mui/material"
import React from "react"


export default function Jumbotron({ title, description, buttonText }) {
  return (
    <Box sx={{ my: 6 }} component={"section"}>
      <Typography variant="h2" component="h2" gutterBottom>
        {title}
      </Typography>
      <Typography paragraph variant="body1">
        {description}
      </Typography>
      <Button variant="contained" color="primary">
        {buttonText}
      </Button>
    </Box>
  )
}
