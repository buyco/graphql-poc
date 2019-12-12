import React, { useState } from 'react';
import { InputGroup, Button } from "@blueprintjs/core";

interface IProps {
  onAdd: {
    (text: string): unknown
  },
}

const AddTodo = ({ onAdd }: IProps) => {
  const [text, setText] = useState("");

  const handleChange = (e: React.FormEvent<HTMLInputElement>) => setText(e.currentTarget.value);

  const handleClick = () => {
    onAdd(text);
    setText("");
  }

  const handleKeyUp = (e: React.KeyboardEvent) => {
    if (e.keyCode === 13) {
      handleClick();
    }
  }

  return (
    <div style={{ padding: "1em 0" }}>
      <InputGroup
        rightElement={
          <Button icon="arrow-right" minimal={true} onClick={handleClick} />
        }
        onChange={handleChange}
        onKeyUp={handleKeyUp}
        placeholder="Add todo"
        value={text}
        large
      />
    </div>
  );
  }

export default AddTodo;