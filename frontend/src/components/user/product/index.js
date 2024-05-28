import React, { useEffect, useState } from 'react';
import Title from './../public/title';
import ProductList from './list';
import ProductModal from './../product/modal'
import ProductDetail from './../product/detail'
import { Container, Tabs, Tab } from 'react-bootstrap';
import 'bootstrap/js/dist/tab'

function Product(props) {

  const { id, keyword, setKeyword, modalId, setModalId } = props
  // get/set data from api display on this component
  const [data, setData] = useState(null)
  // get/set type id and pass into detail and modal component
  const [typeId, setTypeId] = useState(0)
  // get/set type id and pass into detail and modal component
  const [eventIds, setEventIds] = useState([0])
  // get/set prodoct id and pass into detail and modal component
  const [productId, setProductId] = useState(null)

  useEffect(() => {
    fetch(`${process.env.REACT_APP_API_URL}/type/list/2/8`)
    .then((response) => {
      if (!response.ok) {
        console.log('error')
      }
      return response.json()
    })
    .then((json) => {
      setData(json.data)
    })
  }, [])

  const handleClick = (name, id) => {
    let isExist = true
    switch (name) {
      case 'type': 
        setTypeId(id)
        setEventIds([0])
        break
      case 'event': 
        setTypeId(0)
        setEventIds([id, 0])
        break
      default: isExist = false; break;
    }
    if (isExist) setModalId('productList')
  }

  return (
    <React.Fragment>
      <section className='product py-5' id={id}>
        <Container className='container-ctm py-3'>
          <Title setClass='mb-4' title='Type Products' supTitle='Showing latest products of type' />
          <Tabs defaultActiveKey='productTab0' className='product__tab mb-4 pt-2'>
            {data && data.list.map((item, index) => {
              return (
                <Tab eventKey={`productTab` + index} title={item.name} key={index} className='nav-item'>
                  <ProductList setModalId={setModalId} setProductId={setProductId} data={item.products} />
                  <div className='text-center pt-5'>
                    <a className='btn btn-white' href='javascript:;' onClick={() => handleClick('type', item.id)}>More products</a>
                  </div>
                </Tab>
              );
            })}
          </Tabs>
        </Container>
      </section>
      <ProductModal keyword={keyword} setKeyword={setKeyword} setModalId={setModalId} typeId={typeId} setTypeId={setTypeId} eventIds={eventIds} setEventIds={setEventIds} setProductId={setProductId} isShow={`productList` === modalId} />
      <ProductDetail setModalId={setModalId} productId={productId} setProductId={setProductId} isShow={`productDetail` === modalId} handleClick={handleClick} />
    </React.Fragment>
  );
}

export default Product;