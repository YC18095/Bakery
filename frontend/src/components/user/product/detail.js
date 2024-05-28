import React, { useState, useEffect } from 'react';
import RelatedItems from './related-item';
import { Modal, Row, Col, Image, Spinner } from 'react-bootstrap';
import { EffectFade, Pagination } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/scss';
import 'swiper/scss/free-mode';
import 'swiper/scss/thumbs';

function ProductDetail(props) {

  const { setModalId, productId, setProductId, isShow, handleClick } = props
  // get/set data from api display on this component
  const [data, setData] = useState()
  // loading
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    if (productId != null && productId > 0) {
      setLoading(true)
      fetch(`${process.env.REACT_APP_API_URL}/product/detail/${productId}`)
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
    }
  }, [productId])

  return (
    <Modal show={isShow} onHide={() => setModalId('')} scrollable='true' contentClassName='mx-auto product__content' className='product'>
      <Modal.Body className='p-2'>
        {!loading && data ? 
          <React.Fragment>
            <Row className='g-0 pb-2'>
              <Col xs={5}>
                <Swiper
                  className='product__carousel'
                  modules={[ EffectFade, Pagination ]}
                  effect={'fade'}
                  fadeEffect={{ crossFade: true }}
                  pagination={{ clickable: true }}
                >
                  {data.images && data.images.map((image, index) => {
                    return (
                      <SwiperSlide key={index} className='ratio ratio-1x1'>
                        <Image src={`${process.env.REACT_APP_IMG_URL}/${image}`} alt={data.name} />
                      </SwiperSlide>
                    )
                  })}
                </Swiper>
              </Col>
              <Col xs={7} className='p-2 ps-4'>
                <div className='d-flex justify-content-between mb-2'>
                  <a className='icon d-inlie-block' href='javascript:;' onClick={() => setModalId('productList')}>
                    <i className='bi bi-grid'></i>
                  </a>
                  <a className='icon d-inlie-block' href='javascript:;' onClick={() => setModalId('')}>
                    <i className='bi bi-x-lg'></i>
                  </a>
                </div>
                <h2 className='text-truncate mb-3'>{data.name}</h2>
                <div className='product__info'>
                  <Row className='g-0'>
                    <Col xs={2}>Type:</Col>
                    <Col xs={10}>
                      <a className='d-inline-block' href='javascript:;' onClick={() => handleClick('type', data.type.id)}>{data.type.name}</a>
                    </Col>
                  </Row>
                  <Row className='g-0'>
                    <Col xs={2}>Events:</Col>
                    <Col xs={10}>
                      {data.events.map((item, index) => {
                        return (
                          <a key={index} className='d-inline-block' href='javascript:;' onClick={() => handleClick('event', item.id)}>{item.name}</a>
                        )
                      })}
                    </Col>
                  </Row>
                  <Row className='g-0 pt-2 pb-3 border-bottom'>
                    <Col xs={12}>
                      <div className='content'>
                        <p>{data.content}</p>
                      </div>
                    </Col>
                  </Row>
                  <Row className='g-0 pt-2'>
                    <Col xs={2}>Share:</Col>
                    <Col xs={10}>
                      <a className='media d-inline-block' href='#'><i className='bi bi-facebook'></i></a>
                      <a className='media d-inline-block' href='#'><i className='bi bi-twitter'></i></a>
                      <a className='media d-inline-block' href='#'><i className='bi bi-instagram'></i></a>
                      <a className='media d-inline-block' href='#'><i className='bi bi-telegram'></i></a>
                    </Col>
                  </Row>
                </div>
              </Col>
            </Row>
            <Row className='gx-3 pt-2 border-top'>
              {data.relatedItems.map((item, index) => {
                return (
                  <div key={index} className='col-4 m-0 border-end'>
                    <RelatedItems setProductId={setProductId} data={item} />
                  </div>
                );
              })}
            </Row>
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

export default ProductDetail;