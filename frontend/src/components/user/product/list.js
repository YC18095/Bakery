import React from 'react';
import ProductItem from './item';
import { Row, Col } from 'react-bootstrap';

function ProductList(props) {

  const { setModalId, setProductId, data } = props
  
  return (
    <Row>
      {data && data.map((item, index) => {
        return (
          <Col xs={3} key={index} className='mb-2'>
            <ProductItem setModalId={setModalId} setProductId={setProductId} data={item} />
          </Col>
        );
      })}
    </Row>
  );
}

export default ProductList;
