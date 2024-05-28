import React from 'react';
import { Container, Ratio, Image } from 'react-bootstrap';
import { EffectFade, Autoplay, Pagination } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/scss';
import 'swiper/scss/effect-fade';
import 'swiper/scss/pagination';

function Banner() {
  const slides = ['banner-1.jpg', 'banner-2.jpg', 'banner-3.jpg'];
  return (
    <section className='banner'>
      <Container>
        <Swiper
          className='banner__carousel'
          modules={[ EffectFade, Autoplay, Pagination ]}
          effect={'fade'}
          pagination={true}
          autoplay={{ delay: 3000, disableOnInteraction: false }}
          loop={true}
        >
          {slides.map((item, index) => {
            return (
              <SwiperSlide key={index} className='swiper-slide'>
                <Ratio aspectRatio='10x4'>
                  <Image className='d-block w-100' src={`${process.env.REACT_APP_IMG_URL}/${item}`} alt='Slide' />
                </Ratio>
              </SwiperSlide>
            );
          })}
        </Swiper>
      </Container>
    </section>
  );
}

export default Banner;
