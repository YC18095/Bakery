import React, { useState, useEffect} from 'react';
import { Modal, Ratio, Image, Spinner } from 'react-bootstrap';

function NewsDetail(props) {

  const { setModalId, newsId, setNewsId, isShow } = props
  // get/set data from api display on this component
  const [data, setData] = useState(null)
  // loading
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    if (newsId != null && newsId > 0) {
      setLoading(true)
      fetch(`${process.env.REACT_APP_API_URL}/news/detail/${newsId}`)
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
  }, [newsId])

  return (
    <Modal show={isShow} onHide={() => setModalId('')} scrollable='true' contentClassName='mx-auto news__content' className='news'>
      {!loading && data ?
        <React.Fragment>
          <Modal.Header className='py-2 px-3'>
            <div className='w-100 d-flex justify-content-between'>
              <h2 className='mb-0'>{data.name}</h2>
              <a className='icon d-inlie-block' href='javascript:;' onClick={() => setModalId('')}>
                <i className='bi bi-x-lg'></i>
              </a>
            </div>
            <div className='pt-1 tag'>
              <span className='d-inline-block me-3'>
                <i className='bi bi-clock me-1'></i>{new Date(data.createdAt).toLocaleDateString('en', { year: 'numeric', month: 'short', day: 'numeric' })}
              </span>
              <span className='d-inline-block me-3'>
                <i className='bi bi-calendar-event me-1'></i>{data.event.name}
              </span>
            </div>
          </Modal.Header>
          <Modal.Body className='p-3'>
            <Ratio aspectRatio='16x10' className='mb-3'>
              <Image src={`${process.env.REACT_APP_IMG_URL}/${data.image}`} alt={data.name} />
            </Ratio>
            <div className='content pb-2'>
              <p>{data.content}</p>
            </div>
            <div className='share pb-3'>
              <span className='d-inline-block me-3'>Share: </span>
              <a className='media d-inline-block' href='#'>
                <i className='bi bi-facebook'></i>
              </a>
              <a className='media d-inline-block' href='#'>
                <i className='bi bi-twitter'></i>
              </a>
              <a className='media d-inline-block' href='#'>
                <i className='bi bi-instagram'></i>
              </a>
              <a className='media d-inline-block' href='#'>
                <i className='bi bi-telegram'></i>
              </a>
            </div>
            <div className='more border-top pt-3'>
              {data.relatedItems && data.relatedItems.map((item, index) => {
                return (
                  <a key={index} className={`d-block text-truncate'` + (item.id < data.id ? ` mb-1` : null)} href='javascript:;' onClick={() => setNewsId
                    (item.id)}>
                    <span className='me-1'>{item.id < data.id ? `Prev:` : 'Next:'}</span>{item.name}
                  </a>
                )
              })}
            </div>
          </Modal.Body>
        </React.Fragment>
        :
          <div className='loading'>
            <Spinner animation='border' role='status' size='lg' />
          </div>
        }
    </Modal>
  );
}

export default NewsDetail;