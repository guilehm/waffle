import Jumbotron from "@/components/Jumbotron/Jumbotron"
import { Container, Typography } from "@mui/material"


export default function Home() {
  return (
    <main>
      <Container maxWidth="lg">
        <Jumbotron
          title={"Hello World"}
          description={"whatever you want"}
          buttonText={"omg"} />
      </Container>
    </main>
  )
}
