import React from 'react';
import { Row, Col, Ratio, Image } from 'react-bootstrap';

function RelatedItem(props) {

  const { setProductId, data } = props

  return (
    <Row as='a' className='g-0 product__item-snd' href='javascript:;' onClick={() => setProductId(data.id)}>
      <Col xs={4}>
        <Ratio aspectRatio='1x1'>
          <Image src={`${process.env.REACT_APP_IMG_URL}/${data.images[0]}`} alt={data.name} />
        </Ratio>
      </Col>
      <Col xs={8} className='ps-2'>
        <h3 className='mb-0 text-limit-2 text-mh'>{data.name}</h3>
      </Col>
    </Row>
  );
}

export default RelatedItem;