import React from 'react';
import { Modal, Form, Button } from "react-bootstrap";

function ConfirmModal(props) {

  const { modal, setModal, action, dataId } = props

  return (
    <Modal show={modal} onHide={() => setModal(false)} size='md'>
      <React.Fragment>
        <Modal.Header closeButton>
          <Modal.Title>Remove item</Modal.Title>
        </Modal.Header>
        <Modal.Body>Are you sure to remove this item?</Modal.Body>
        <Modal.Footer>
          <Form method='post' action={action} className='w-100 text-center'>
            <Form.Control type='hidden' name='id' value={dataId} />
            <Button variant='primary' type='submit' className='me-3'>Sure</Button>
            <Button variant='secondary' onClick={() => setModal(false)}>Cancel</Button>
          </Form>
        </Modal.Footer>
      </React.Fragment>
    </Modal>
  );
}

export default ConfirmModal;