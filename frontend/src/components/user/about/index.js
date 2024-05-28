import React from 'react';
import { Container, Row, Col, Ratio, Image } from 'react-bootstrap';
import Title from './../public/title';

function About(props) {

  const { id } = props

  const data = {
    title: 'Bakery',
    background: 'about-bg.jpg',
    images: ['about-store-1.jpg', 'about-store-2.jpg'],
    content: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.'
  }

  return (
    <section className='about' id={id}>
      <Container className='border-top pt-5'>
        <Title setClass='mb-5' title='About Us' supTitle='The history of business' />
        <div className='bgi-cover about__bg p-5' style={{ backgroundImage: `url(${process.env.REACT_APP_IMG_URL}/${data.background})`}}>
          <Row className='g-0'>
            <Col xs={6} className='offset-6 about__box'>
              <h6 className='mb-2'><span className='d-inline-block pb-1'>{data.title}</span></h6>
              <div className='content pt-1'>
                <p>{data.content}</p>
              </div>
              <Row className='gx-3 pt-4'>
                {data.images.map((item, index) => {
                  return (
                    <Col xs={6} key={index}>
                      <Ratio aspectRatio='16x10'>
                        <Image src={`${process.env.REACT_APP_IMG_URL}/${item}`} alt='' />
                      </Ratio>
                    </Col>
                  );
                })}
              </Row>
            </Col>
          </Row>
        </div>
      </Container>
    </section>
  );
}

export default About;