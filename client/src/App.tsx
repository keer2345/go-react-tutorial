import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/custom/Navbar";
import TodoForm from "./components/custom/TodoForm";

function App() {
  return (
    <>
      <Stack h="100vh">
        <Navbar />
        <Container>
          <TodoForm />
          {/* <TodoList /> */}
        </Container>
      </Stack>
    </>
  );
}

export default App;
