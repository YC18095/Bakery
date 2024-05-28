import React, { useState, useEffect } from 'react';
import { Nav, Image } from 'react-bootstrap';
import { Link } from 'react-router-dom'

function Sidebar() {

  const [ active, setActive ] = useState(0)

  const menu = [
    { name: 'Dashboard', href: 'dashboard', icon: 'bi-pie-chart-fill' },
    { name: 'Types', href: 'type', icon: 'bi-grid-1x2-fill' },
    { name: 'Events', href: 'event', icon: 'bi-calendar-fill' },
    { name: 'Products', href: 'product', icon: 'bi-grid-fill' },
    { name: 'News', href: 'news', icon: 'bi-newspaper' },
  ]

  useEffect(() => {
    const pathname = window.location.pathname.split('/');
    menu.forEach((item, index) => {
      if (item.href === pathname[2]) {
        setActive(index)
        return
      }
    });
  }, [])

  return (
    <div className='sidebar py-3 pe-4'>
      <div className='sidebar__box'>
        <a className='d-block sidebar__logo py-1' href='#'>
          <Image src='../../../images/logo-none.png' />
        </a>
        <Nav className='sidebar__navbar py-2 px-3'>
          {menu.map((item, index) => {
            return (
              <Link key={index} className={`nav-link mb-1${active === index ? ' active' : ''}`} to={item.href} onClick={() => setActive(index)}>
                <i className={`bi ${item.icon} me-2`}></i>{item.name}
              </Link>
            );
          })}
        </Nav>
      </div>
    </div>
  );
}

export default Sidebar;