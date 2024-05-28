import React, { useState } from 'react';
import Header from './header'
import Banner from './banner'
import Product from './product'
import About from './about'
import News from './news'
import Footer from './footer'
import 'bootstrap/js/dist/scrollspy'

function UserLayout() {

  const modalIdList = ['productList', 'productDetail', 'newsList', 'newsDetail']

  const menu = {
    id: 'headerNavbar',
    list: [
      { name: 'Product', href: 'section0' },
      { name: 'About', href: 'section1' },
      { name: 'Events', href: 'section2' },
      {
        name: 'List', href: 'javascript:;', list: [
          { name: 'Product List', href: modalIdList[0] },
          { name: 'Events List', href: modalIdList[2] },
        ]
      },
    ]
  }

  // eslint-disable-next-line
  const [keyword, setKeyword] = useState(null)
  // const [productListModal, setProductListModal] = useState(false)
  const [modalId, setModalId] = useState('')

  return (
    <div data-bs-spy='scroll' data-bs-target={`#` + menu.id}>
      <Header menu={menu} keyword={keyword} setKeyword={setKeyword} setModalId={setModalId} />
      <Banner />
      <Product id={menu.list[0].href} keyword={keyword} setKeyword={setKeyword} modalId={modalId} setModalId={setModalId} />
      <About id={menu.list[1].href} />
      <News id={menu.list[2].href} modalId={modalId} setModalId={setModalId} />
      <Footer />
    </div>
  );
}

export default UserLayout