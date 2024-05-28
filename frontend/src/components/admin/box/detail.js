import React from 'react';
import Field from './field';
import BadgeStatus from '../public/badge-status';

function Detail(props) {

  const { data } = props

  return (
    <div className='admin__detail'>
      {data.name ? <Field name='Name'>{data.name}</Field> : ''}
      {data.status ? <Field name='Status'><BadgeStatus status={data.status} /></Field> : ''}
      {data.createdAt ? <Field name='Created at'>{new Date(data.createdAt).toLocaleDateString('en', { year: 'numeric', month: 'short', day: 'numeric' })}</Field> : ''}
      {data.updatedAt ? <Field name='Last update'>{new Date(data.updatedAt).toLocaleDateString('en', { year: 'numeric', month: 'short', day: 'numeric' })}</Field> : ''}
    </div>
  );
}

export default Detail;