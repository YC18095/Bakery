import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Row, Col } from 'react-bootstrap';
import Box from '../box';
import Relation from './relation';

function TypeEdit() {

  const { id } = useParams();

  const [data, setData] = useState(null)

  useEffect(() => {
    fetch('http://localhost:8000/api/admin/type/detail/' + id)
      .then(response => {
        if (!response.ok) {
          console.log('error')
        }
        return response.json()
      })
      .then(json => {
        setData(json.data)
      })
  }, [])

  return (
    <Row>
      {data ?
        <React.Fragment>
          <Col xs={8}>
            <Box>
              <Box.Title title='Type Details' />
              <Box.Table action={`http://localhost:8000/api/admin/type/edit`} data={data} />
            </Box>
          </Col>
          <Col xs={4}>
            <Box>
              <Relation typeId={id} location='edit' />
            </Box>
          </Col>
        </React.Fragment>
        : ''}
    </Row>
  );
}

export default TypeEdit;