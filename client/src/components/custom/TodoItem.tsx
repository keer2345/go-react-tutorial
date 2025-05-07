import { BASE_URL } from "@/App";
import { Todo } from "./TodoList";
import { Badge, Box, Flex, Text } from "@chakra-ui/react";
import axios from "axios";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { mutate } from "swr";

const TodoItem = ({ todo }: { todo: Todo }) => {
  const updateTodo = async () => {
    try {
      const updateTodo = await axios({
        method: "PATCH",
        url: BASE_URL + `/todo/+${todo._id}`,
      }).then((res) => {
        console.log("v5:", res);
        return res.data;
      });

      mutateTodo(updateTodo, false);
    } catch (error) {
      console.log("update failed", error);
    }
  };
  return (
    <Flex gap={2} alignItems={"center"}>
      <Flex
        flex={1}
        alignItems={"center"}
        border={"1px"}
        borderColor={"gray.600"}
        p={2}
        borderRadius={"lg"}
        justifyContent={"space-between"}
      >
        <Text
          color={todo.completed ? "green.200" : "yellow.100"}
          textDecoration={todo.completed ? "line-through" : "none"}
        >
          {todo.body}
        </Text>
        {todo.completed && (
          <Badge ml="1" colorScheme="green">
            Done
          </Badge>
        )}
        {!todo.completed && (
          <Badge ml="1" colorScheme="yellow">
            In Progress
          </Badge>
        )}
      </Flex>
      <Flex gap={2} alignItems={"center"}>
        <Box
          color={"green.500"}
          cursor={"pointer"}
          onClick={() => updateTodo()}
        >
          <FaCheckCircle size={20} />
        </Box>
        <Box
          color={"red.500"}
          cursor={"pointer"}
          onClick={() => console.log("b")}
        >
          <MdDelete size={25} />
        </Box>
      </Flex>
    </Flex>
  );
};

export default TodoItem;
