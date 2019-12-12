

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: getTodos
// ====================================================

export interface getTodos_todos_user {
  name: string;
}

export interface getTodos_todos {
  id: string;
  text: string;
  user: getTodos_todos_user;
}

export interface getTodos {
  todos: getTodos_todos[];
}


/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: createTodo
// ====================================================

export interface createTodo_createTodo_user {
  name: string;
}

export interface createTodo_createTodo {
  user: createTodo_createTodo_user;
  text: string;
}

export interface createTodo {
  createTodo: createTodo_createTodo;
}

export interface createTodoVariables {
  userId: string;
  text: string;
}


/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: deleteTodo
// ====================================================

export interface deleteTodo {
  deleteTodo: boolean;
}

export interface deleteTodoVariables {
  todoId: string;
}

/* tslint:disable */
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

//==============================================================
// END Enums and Input Objects
//==============================================================