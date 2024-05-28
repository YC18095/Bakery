import React from 'react';
import Menu from './menu';
import Search from '../public/search';
import { Container, Image } from 'react-bootstrap';
import 'bootstrap/js/dist/modal'

function Header(props) {

  const { menu, keyword, setKeyword, setModalId } = props

  const logo = 'logo.png';
  
  return (
    <header className='header'>
      <Container>
        <div className='d-table w-100'>
          <div className='d-table-cell header__logo'>
            <a className='d-inline-block' href='index.html'>
              <Image src={`${process.env.REACT_APP_IMG_URL}/${logo}`} alt='Bakery' />
            </a>
          </div>
          <Menu menu={menu} setModalId={setModalId} />
          <div className='d-table-cell header__feature'>
            <Search className='w-100 position-relative' keyword={keyword} setKeyword={setKeyword} setModalId={setModalId} isHeader={true} />
          </div>
        </div>
      </Container>
    </header>
  );
}

export default Header;