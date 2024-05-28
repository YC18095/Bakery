import React, { useEffect, useState } from 'react';
import NewsItem from './item';
import Page from '../public/page';
import { Modal, Row, Col, Spinner } from 'react-bootstrap';

function NewsList(props) {

  const { setModalId, setNewsId, isShow } = props
  // get/set data from api display on this component
  const [data, setData] = useState(null)
  // get/set type id from api nad pass into feature component
  const [page, setPage] = useState(0)
  // loading
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    fetch(`${process.env.REACT_APP_API_URL}/news/list/6/${page}`)
      .then((response) => {
        if (!response.ok) {
          console.log('error')
        }
        return response.json()
      })
      .then((json) => {
        setData(json.data)
        setLoading(false)
      })
  }, [page])

  return (
    <Modal show={isShow} onHide={() => setModalId('')} scrollable='true' contentClassName='mx-auto news__list' className='news'>
      <Modal.Header className='px-3 py-2'>
        <h6 className='mb-0'>News List</h6>
        <a className='icon d-inlie-block' href='javascript:;' onClick={() => setModalId('')}>
          <i className='bi bi-x-lg'></i>
        </a>
      </Modal.Header>
      <Modal.Body className='p-3'>
        {!loading && data ? 
          <React.Fragment>
            <Row className='gx-3'>
              {data.list.map((item, index) => {
                return (
                  <Col xs={4} key={index} className='mb-3'>
                    <NewsItem setModalId={setModalId} setNewsId={setNewsId} data={item} />
                  </Col>
                );
              })}
            </Row>
            <Page page={page} setPage={setPage} total={data ? Math.ceil(data.total / 6) : 0} className='justify-content-center' />
          </React.Fragment>
        :
          <div className='loading'>
            <Spinner animation='border' role='status' size='lg' />
          </div>
        }
      </Modal.Body>
    </Modal>
  );
}

export default NewsList;