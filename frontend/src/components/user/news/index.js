import React, { useState, useEffect } from 'react';
import Title from './../public/title';
import NewsList from './../news/list'
import NewsDetail from './../news/detail'
import NewsItem from './item';
import { Navigation } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/scss';

function News(props) {

  const { id, modalId, setModalId } = props
  // get/set data from api display on this component
  const [data, setData] = useState(null)
  // get/set news id and pass into detail and modal component
  const [newsId, setNewsId] = useState(null)

  useEffect(() => {
    fetch(`${process.env.REACT_APP_API_URL}/news/list/5/0`)
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

  return (
    <React.Fragment>
      <section className='news py-5' id={id}>
        <div className='container pt-4'>
          <Title setClass='mb-5' title='Event`s Story' supTitle='Talking about the festival`s origin' />
          <Swiper
            className='news__carousel'
            modules={[ Navigation ]}
            slidesPerView={3}
            spaceBetween={20}
            navigation={{ prevEl: '.news__carousel .swiper-button-prev', nextEl: '.news__carousel .swiper-button-next' }}
          >
            {data && data.list.map((item, index) => {
              return (
                <SwiperSlide key={index}>
                  <NewsItem setModalId={setModalId} setNewsId={setNewsId} data={item} />
                </SwiperSlide>
              );
            })}
            <div className='swiper-button-prev'><i className='bi bi-chevron-left'></i></div>
            <div className='swiper-button-next'><i className='bi bi-chevron-right'></i></div>
          </Swiper>
        </div>
      </section>
      <NewsList setModalId={setModalId} setNewsId={setNewsId} isShow={`newsList` === modalId} />
      <NewsDetail setModalId={setModalId} newsId={newsId} setNewsId={setNewsId} isShow={`newsDetail` === modalId} />
    </React.Fragment>
  );
}

export default News;
