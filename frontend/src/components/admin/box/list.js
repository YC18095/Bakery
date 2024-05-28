import React from 'react';
import { Table } from "react-bootstrap";

function List(props) {

  const { fields, children } = props

  return (
    <Table responsive hover className='admin__list'>
      <thead>
        <tr>
          {fields.map((item, index) => {
            return (
              <td key={index} className={item.className}>{item.name}</td>
            );
          })}
        </tr>
      </thead>
      <tbody>
        {children}
      </tbody>
    </Table>
  );
}

export default List;