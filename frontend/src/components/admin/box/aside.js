import React from 'react';

function Aside(props) {

  const { title, type, page, setPage, total, children } = props

  return (
    <div className='pt-4 admin__relation'>
      <h6 className='mb-3 text-center'>{title}</h6>
      <div className='d-flex justify-content-between align-items-center mb-2'>
        <div className='title'>{type}</div>
        {total > 1 ? 
          <div className='page'>
            {page > 0 ? 
              <a href='javascript:;' onClick={() => setPage(page-1)}>
                <i className="bi bi-chevron-left me-1"></i>
              </a>
            : ''}
            <span>{page+1}</span>
            {page < total-1 ? 
              <a href='javascript:;' onClick={() => setPage(page+1)}>
                <i className="bi bi-chevron-right ms-1"></i>
              </a>
            : ''}
          </div>
        : ''}
      </div>
      {children}
    </div>
  );
}

export default Aside;