import React from 'react';
import { Form, Button } from "react-bootstrap";
import Field from './field';

function Table(props) {

  const { action, data } = props

  return (
    <div className='admin__form'>
      <Form method='post' action={action}>
        <Form.Control type='hidden' name='id' value={data.id} />
        <Field name='Name'>
          <Form.Control type='text' name='name' defaultValue={data.name} />
          <Form.Text className="text-muted">
            We'll never share your email with anyone else.
          </Form.Text>
        </Field>
        <Field name='Status'>
          {[{ key: 'Active', value: 1 }, { key: 'Unactive', value: 2 }].map((item, index) => {
            return (
              <Form.Check id={`status${index}`} className='d-inline-block me-4' type='radio' name='status' defaultValue={item.value} label={item.key} defaultChecked={item.value === data.status} />
            );
          })}
        </Field>
        <Field>
          <Button type='submit' variant='primary' className='me-3'><i className="bi bi-upload me-1"></i> Submit</Button>
          <Button type='button' variant='secondary' onClick={() => {window.history.back()}}><i className="bi bi-arrow-90deg-left me-1"></i>Cancel</Button>
        </Field>
      </Form>
    </div>
  );
}

export default Table;