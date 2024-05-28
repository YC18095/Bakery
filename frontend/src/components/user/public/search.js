import React from 'react';

function Search(props) {
  const { className, keyword, setKeyword, setModalId, isHeader } = props

  const handleKeyDown = event => {
    if (event.key === 'Enter' && isHeader) setModalId('productList')
  }

  return (
    <div className={`form-search` + (className ? ` ` + className : ``)}>
      <input type='text' placeholder='Search...' value={keyword} onChange={event => setKeyword(event.target.value)} onKeyDown={event => handleKeyDown(event)} />
      <button type='button'><i className='bi bi-search'></i></button>
    </div>
  );
}

export default Search;