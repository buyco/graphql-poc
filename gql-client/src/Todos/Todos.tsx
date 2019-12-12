import React from "react";
import { gql } from "apollo-boost";
import { useQuery, useMutation } from "@apollo/react-hooks";
import {
  createTodo, // eslint-disable-line
  getTodos,
  deleteTodo // eslint-disable-line
} from "../../operation-result-types";
import TodoList from "./TodoList";
import AddTodo from "./AddTodo";

const GET_TODOS = gql`
  query getTodos {
    todos {
      id
      text
      user {
        name
      }
    }
  }
`;

const CREATE_TODO = gql`
  mutation createTodo($userId: ID!, $text: String!) {
    createTodo(input: { text: $text, userId: $userId }) {
      user {
        name
      }
      text
      id
    }
  }
`;

const DELETE_TODO = gql`
  mutation deleteTodo($todoId: ID!) {
    deleteTodo(input: { id: $todoId })
  }
`;

const Todos = () => {
  const { loading, data, refetch } = useQuery<getTodos>(GET_TODOS);

  const [deleteTodo] = useMutation<deleteTodo>(DELETE_TODO, {
    update() {
      refetch();
    }
  });

  const [createTodo] = useMutation<createTodo>(CREATE_TODO, {
    update(cache, { data }) {
      if (data) {
        const { todos } = cache.readQuery({ query: GET_TODOS }) || {};

        cache.writeQuery({
          query: GET_TODOS,
          data: { todos: [...todos, data.createTodo] }
        });
      }
    }
  });

  const handleAddClick = (text: string) =>
    createTodo({
      variables: { userId: "T5577006791947779410", text }
    });

  const handleDeleteClick = (todoId: string) =>
    deleteTodo({
      variables: { todoId }
    });

  return (
    <>
      <AddTodo onAdd={handleAddClick} />
      <TodoList
        loading={loading}
        todos={data && data.todos}
        deleteTodo={handleDeleteClick}
      />
    </>
  );
};

export default Todos;
