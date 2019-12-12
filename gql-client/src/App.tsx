import React from 'react';
import '@blueprintjs/core/lib/css/blueprint.css';
import 'normalize.css/normalize.css';
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Todos from "./Todos/Todos";
import "./App.scss";

const client = new ApolloClient({
  uri: "http://localhost:8080/query"
});

const App: React.FC = () => {
  return (
    <ApolloProvider client={client}>
      <div className="app bp3-dark">
          <Todos />
      </div>
    </ApolloProvider>
  );
}

export default App;
