import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: JSX.Element;
};


// todo x: 首页内容
const FeatureList: FeatureItem[] = [
  {
    title: 'ZeroMall: B2B2C 电商平台',
    Svg: require('@site/static/img/go.svg').default,
    description: (
      <>
        <ul>
          <li>基于: Go + go-zero + gRPC 微服务开发.</li>
          <li>集成中间件: mysql, redis, kafka, consul 等.</li>
          <li>API: http rest/grpc/websocket.</li>
        </ul>
      </>
    ),
  },
  {
    title: 'eShop: S2B2C 电商平台',
    Svg: require('@site/static/img/logo.svg').default,
    description: (
      <>
        <ul>
          <li>eShop: S2B2C 电商平台.</li>
          <li>集成授权码/license 售卖服务.</li>
          <li>集成微信/支付宝/数字货币等支付工具.</li>
        </ul>
      </>
    ),
  },


  {
    title: 'LiveMall: 直播电商平台',
    Svg: require('@site/static/img/go.svg').default,
    description: (
      <>
        <ul>
          <li>LiveMall: 直播电商平台</li>
          <li>Live/Vlog: 直播/短视频服务</li>
          <li>Mall: 电商服务</li>
        </ul>
      </>
    ),
  },


];

function Feature({title, Svg, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img"/>
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
