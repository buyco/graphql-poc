import React from "react";
import { Spinner, Card, Button } from "@blueprintjs/core";
import { getTodos_todos } from "../../operation-result-types";

interface IProps {
  todos: Array<getTodos_todos> | undefined,
  loading: Boolean,
  deleteTodo: {
    (todoId: string): void
  }
}

const TodoList = ({ todos, deleteTodo, loading }: IProps) => {
  if (loading) return <Spinner />;

  if (!todos) return null;

  return (
    <div>
      {todos.map(todo => (
        <Card>
          {todo.text}
          <Button style={{ float: 'right' }} icon="delete" onClick={() => deleteTodo(todo.id)} minimal />
        </Card>
      ))}
    </div>
  );
};

export default TodoList;
