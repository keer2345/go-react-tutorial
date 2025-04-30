import { Button, Flex, Input } from "@chakra-ui/react";
import React, { useState } from "react";

const TodoForm = () => {
  const {} = useMutation({});
  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input />
        <Button>aa</Button>
      </Flex>
    </form>
  );
};

export default TodoForm;
