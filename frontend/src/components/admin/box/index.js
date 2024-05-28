import React from 'react';
import Title from './title';
import List from './list';
import Detail from './detail';
import Table from './table';
import Field from './field';
import Aside from './aside';

function Box(props) {

  const { children } = props

  return (
    <div className='admin__box px-3 pb-3'>
      {children}
    </div>
  );
}

Box.Title = Title;
Box.List = List;
Box.Detail = Detail;
Box.Table = Table;
Box.Field = Field;
Box.Aside = Aside;
export default Box;