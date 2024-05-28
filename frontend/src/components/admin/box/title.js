import React from 'react';
import { Link } from "react-router-dom";
import { Button } from "react-bootstrap";

function Title(props) {

  const { title, create, edit, setModal } = props

  return (
    <div className='admin__title px-3 py-4'>
      <span>{title}</span>
      <div className='admin__button'>
        {create ? <Link to={create}><i className='bi bi-plus-lg'></i>Create</Link> : ''}
        {edit ? <Link to={edit}><i className='bi bi-pen-fill'></i>Edit</Link> : ''}
        {setModal ? <a className='ms-3' href='javascript:;' onClick={() => setModal(true)}><i className="bi bi-trash-fill"></i>Remove</a> : ''}
      </div>
    </div>
  );
}

export default Title;