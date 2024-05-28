import React from 'react';
import { Ratio, Image } from 'react-bootstrap';
import 'bootstrap/js/dist/modal'

function NewsItem(props) {
  
  const { setModalId, setNewsId, data } = props

  let date = new Date(data.createdAt)
  
  const handleClick = event => {
    let thisElement = event.target.closest('.news__item')
    setNewsId(thisElement.dataset.id)
    setModalId('newsDetail')
  }

  return (
    <a className='news__item' href='javascript:;' data-id={data.id} onClick={event => handleClick(event)}>
      <span className='d-inline-block date'>{date.getDate()}<span className='d-block'>{date.toLocaleString('default', { month: 'short' })}</span></span>
      <Ratio aspectRatio='16x10'>
        <Image src={`${process.env.REACT_APP_IMG_URL}/${data.image}`} alt={data.name} />
      </Ratio>
      <div className='p-2'>
        <h3 className='text-trnucate mb-1'>{data.name}</h3>
        <div className='text-limit-2 text-h'>
          <div className='content'>
            <p>{data.content}</p>
          </div>
        </div>
      </div>
    </a>
  );
}

export default NewsItem;