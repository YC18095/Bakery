import React, { useEffect, useState } from 'react';
import ProductFeature from './feature';
import ProductList from './list';
import Page from '../public/page';
import { Modal, Row, Col, Spinner, Alert } from 'react-bootstrap';

function ProductModal(props) {

  const { keyword, setKeyword, setModalId, typeId, setTypeId, eventIds, setEventIds, setProductId, isShow } = props
  // get/set data from api display on this component
  const [data, setData] = useState(null)
  // get/set type id from api nad pass into feature component
  const [page, setPage] = useState(0)

  const [currentPage, setCurrentPage] = useState(page)
  // loading
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    if (currentPage === page) setPage(0)
    let url = `${process.env.REACT_APP_API_URL}/product/list/${typeId}/${eventIds}/12/${page}`
    if (keyword != null && keyword != '') url += `/${keyword}`
    fetch(url)
      .then((response) => {
        if (!response.ok) {
          console.log('error')
        }
        return response.json()
      })
      .then((json) => {
        setCurrentPage(page)
        setData(json.data)
        setLoading(false)
      })
  }, [keyword, typeId, eventIds, page])

  return (
    <Modal show={isShow} onHide={() => setModalId('')} scrollable='true' contentClassName='mx-auto product__list' className='product'>
      <Modal.Header className='px-3 py-2'>
        <h6 className='mb-0'>Product List</h6>
        <a className='icon d-inlie-block' href='javascript:;' onClick={() => setModalId('')}>
          <i className='bi bi-x-lg'></i>
        </a>
      </Modal.Header>
      <Modal.Body>
        <Row className='h-100'>
          <Col xs={3} className='border-end'>
            <ProductFeature keyword={keyword} setKeyword={setKeyword} typeId={typeId} setTypeId={setTypeId} eventIds={eventIds} setEventIds={setEventIds} />
          </Col>
          <Col xs={9} className='position-relative'>
            {!loading && data ? 
              (data.total > 0 ?
                <React.Fragment>
                  <ProductList setModalId={setModalId} setProductId={setProductId} data={data ? data.list : null} />
                  <Page page={page} setPage={setPage} total={data ? Math.ceil(data.total / 12) : 0} className='justify-content-end' />
                </React.Fragment>
                : <Alert variant='danger'>There're have no any items with <b>{keyword}</b> keywork!</Alert>)
            :
              <div className='loading'>
                <Spinner animation='border' role='status' size='lg' />
              </div>
            }
          </Col>
        </Row>
      </Modal.Body>
    </Modal>
  );
}

export default ProductModal;
