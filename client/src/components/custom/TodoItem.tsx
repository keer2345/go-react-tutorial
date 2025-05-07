import React from "react";
import { Todo } from "./TodoList";

const TodoItem = ({ todo }: { todo: Todo }) => {
  return <div>{todo.body}</div>;
};

export default TodoItem;
