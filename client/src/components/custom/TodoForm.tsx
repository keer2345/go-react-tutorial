/* eslint-disable @typescript-eslint/no-explicit-any */

import useSWR from "swr";

// https://github.com/bezkoder/react-query-axios-typescript
// https://www.bezkoder.com/react-query-axios-typescript/

const fetcher = (url) => fetch(url).then((res) => res.json());

const TodoForm = () => {
  const { data, error, isLoading } = useSWR(
    "http://localhost:5000/api/todo",
    fetcher
  );

  if (error) return <div>Failed to load</div>;
  if (isLoading) return <div>Loading...</div>;

  return <div>TodoForm </div>;
};

export default TodoForm;
