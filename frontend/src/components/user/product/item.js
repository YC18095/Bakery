import React from 'react';
import { Ratio, Image } from 'react-bootstrap';
import 'bootstrap/js/dist/modal'

function ProductItem(props) {

  const { setModalId, setProductId, data } = props
  
  const handleClick = event => {
    let thisElement = event.target.closest('.product__item')
    setProductId(thisElement.dataset.id)
    setModalId('productDetail')
  }

  return (
    // eslint-disable-next-line
    <a className='product__item px-2' href='javascript:;' data-id={data.id} onClick={event => handleClick(event)}>
      <Ratio aspectRatio='1x1' className='w-75 mx-auto'>
        <Image src={`${process.env.REACT_APP_IMG_URL}/${data.images[0]}`} alt={data.name} />
      </Ratio>
      <h3 className='mb-1 text-truncate text-center'>{data.name}</h3>
      <div className='text-limit-2 text-h text-center'>
        <div className='content'>
          <p>{data.content}</p>
        </div>
      </div>
    </a>
  );
}

export default ProductItem;