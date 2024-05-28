import React from 'react';

function Title(props) {
  const { setClass, title, supTitle } = props
  return (
    <div className={`title` + (setClass ? ` ` + setClass : ``)}>
      <span className='d-block text-center mb-1'>{supTitle}</span>
      <h6 className='text-center mb-0'>{title}</h6>
    </div>
  );
}

export default Title;