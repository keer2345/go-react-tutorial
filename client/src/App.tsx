import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/custom/Navbar";
import TodoForm from "./components/custom/TodoForm";
import TodoList from "./components/custom/TodoList";

export const BASE_URL =
  import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";

function App() {
  return (
    <>
      <Stack h="100vh">
        <Navbar />
        <Container>
          {/* <TodoForm /> */}
          <TodoList />
        </Container>
      </Stack>
    </>
  );
}

export default App;
