import React from 'react';
import { Row, Col } from "react-bootstrap";

function Field(props) {

  const { name, children } = props

  return (
    <Row className='g-0 py-3 px-2'>
      <Col xs={3}>{name}:</Col>
      <Col xs={9}>{children}</Col>
    </Row>
  );
}

export default Field;