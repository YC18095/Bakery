import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Row, Col } from 'react-bootstrap';
import Box from '../box';
import Relation from './relation';
import ConfirmModal from '../public/confirm-modal';


function TypeDetail() {

  const { id } = useParams();

  const [ type, setType ] = useState(null)

  const [ modal, setModal ] = useState(false)

  useEffect(() => {
    fetch('http://localhost:8000/api/admin/type/detail/' + id)
      .then(response => {
        if (!response.ok) {
          console.log('error')
        }
        return response.json()
      })
      .then(json => {
        setType(json.data)
      })
  }, [])

  return (
    <React.Fragment>
      <Row>
        {type ? 
          <React.Fragment>
            <Col xs={8}>
              <Box>
                <Box.Title title='Type Details' edit={`/admin/type/edit/${id}`} setModal={setModal} />
                <Box.Detail data={type} />
              </Box>
            </Col>
            <Col xs={4}>
              <Box>
                <Relation typeId={id} />
              </Box>
            </Col>
          </React.Fragment>
        : ''}
      </Row>
      <ConfirmModal modal={modal} setModal={setModal} dataId={id} action='http://localhost:8000/api/admin/type/remove' />
    </React.Fragment>
  );
}

export default TypeDetail;