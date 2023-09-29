import { useState } from "react";

function ListGroup() {
  let items = ["New York", "San Francisco", "Tokyo", "London", "Paris"];

  const [selectedIndex, setSelectedIndex] = useState(-1);

  // Event Handler

  return (
    <>
      <h1>List Group</h1>
      <ul className="list-group">
        {items.length === 0 && <p>No item found</p>}
        {items.map((item, index) => (
          <li
            key={item}
            className={
              selectedIndex === index
                ? "list-group-item active"
                : "list-group-item"
            }
            onClick={() => {
              setSelectedIndex(index);
            }}
          >
            {item}
          </li>
        ))}
      </ul>
    </>
  );
}

export default ListGroup;
