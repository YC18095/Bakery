import React from 'react';
import { Pagination } from 'react-bootstrap';

function Page(props) {
  const { page, setPage, total, className } = props

  let pageArray = []
  for (let i = 0; i < total; i++) {
    pageArray.push(i)
  }

  const clickPage = numsPage => {
    setPage(numsPage)
  }

  return (
    <Pagination className={`py-3` + (className ? ` ` + className : null)}>
      {pageArray.map(index => {
        return (
          <Pagination.Item active={page === index} key={index} onClick={() => clickPage(index)}>{index + 1}</Pagination.Item>
        );
      })}
    </Pagination>
  );
}

export default Page;