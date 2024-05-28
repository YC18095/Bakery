import React from 'react';
import { Nav } from 'react-bootstrap';

function Menu(props) {

  const { menu, setModalId } = props

  return (
    <div className='d-table-cell header__navbar' id={menu.id}>
      <Nav className='justify-content-center'>
        {menu.list.map((item, index) => {
          return (
            <Nav.Item key={index}>
              <Nav.Link href={`#` + item.href}>{item.name}{item.list ? <i className='bi bi-caret-down-fill ms-1'></i>: ``}</Nav.Link>
              {item.list ? 
                <Nav className='header__list py-2'>
                  {item.list.map((subItem, subIndex) => {
                    return (
                      <Nav.Item key={subIndex}>
                        <Nav.Link href='javascript:;' onClick={() => setModalId(subItem.href)}>{subItem.name}</Nav.Link>
                      </Nav.Item>
                    );
                  })}
                </Nav>
              : ``}
            </Nav.Item>
          );
        })}
      </Nav>
    </div>
  );
}

export default Menu;