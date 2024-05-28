import React from 'react';
import { Container } from 'react-bootstrap';

function Footer() {

  return (
    <footer className='footer img-cover'>
      <div className='border-bottom border-light py-4'>
        <Container>
          <div className='footer__info'>
            <div className='mb-2'>
              <a className='d-inline-block' href='tel:0903876321'><i className='bi bi-envelope-fill me-2'></i>0903876321</a>
            </div>
            <div className='mb-2'>
              <a className='d-inline-block' href='mailto:hongky@gmail.com'><i className='bi bi-telephone-fill me-2'></i>hongky@gmail.com</a>
            </div>
            <div>
              <span className='d-inline-block'><i className='bi bi-geo-alt-fill me-2'></i>13 Pho Co Dieu Street, Ward 6, District 11, HCM city</span>
            </div>
          </div>
        </Container>
      </div>
      <div className='py-3'>
        <Container><span className='d-inline-block'>&copy; 2022 | Created by Bakery</span></Container>
      </div>
    </footer>
  );
}

export default Footer;
