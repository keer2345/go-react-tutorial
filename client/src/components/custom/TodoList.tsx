import React from "react";
import { BASE_URL } from "@/App";
import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import useSWR from "swr";
import axios from "axios";
import TodoItem from "./TodoItem";

export type Todo = {
  _id: number;
  body: string;
  completed: boolean;
};

// const fetcher = (url) => fetch(url).then((r) => r.json());
const fetcher = (url: string) => axios.get(url).then((res) => res.data);

const TodoList = () => {
  const {
    data: todos,
    error,
    isLoading,
  } = useSWR<Todo[]>(BASE_URL + "/todo", fetcher);

  return (
    <>
      <Text
        fontSize={"4xl"}
        textTransform={"uppercase"}
        fontWeight={"bold"}
        textAlign={"center"}
        my={2}
        bgGradient="linear(to-l, #0b85f8, #00ffff)"
      >
        Today's Tasks
      </Text>
      {error && (
        <Text alignItems={"center"} color={"red"}>
          {error.message}
        </Text>
      )}
      {!error && (
        <>
          {isLoading && (
            <Flex justifyContent={"center"} my={4}>
              <Spinner size={"xl"} />
            </Flex>
          )}
          {!isLoading && todos?.length == 0 && (
            <Stack alignItems={"center"} gap="3">
              <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
                All tasks completed! 🤞
              </Text>
              <img src="/go.png" alt="Go logo" width={70} height={70} />
            </Stack>
          )}
          <Stack gap={3}>
            {todos?.map((todo) => <TodoItem key={todo._id} todo={todo} />)}
          </Stack>
        </>
      )}
    </>
  );
};

export default TodoList;
