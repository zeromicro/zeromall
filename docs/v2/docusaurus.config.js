// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "ZeroMall",
  tagline: "a Mall base on Go + Go-Zero.",
  url: "https://zeromicro.github.io/zeromall/", // todo x: update
  baseUrl: "/zeromall/",  // todo x: fix for github pages
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.ico",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "zeromicro", // todo x: Usually your GitHub org/user name.
  projectName: "zeromall", // todo x: Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "zh-CN",
    locales: ["zh-CN", "en", "fr"],

    // localeConfigs: {
    //     "zh-CN": {
    //         label: '中文(简体)',
    //         direction: 'ltr',
    //         htmlLang: 'zh-CN',
    //         calendar: 'gregory',
    //         path: 'zh-CN',
    //     },
    //
    //     fr: {
    //         label: 'French',
    //         direction: 'ltr',
    //         htmlLang: 'fr',
    //         calendar: 'gregory',
    //         path: 'fr',
    //     },
    // },


  },

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve("./sidebars.js"),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/",
        },

        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/",
        },

        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      }),
    ],
  ],

  themeConfig:
  /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: "ZeroMall", // todo x
        logo: {
          alt: "Zero Logo",
          src: "img/logo.svg",
        },
        items: [
          // TODO X: 教程
          // {
          //   type: "doc",
          //   docId: "intro",
          //   position: "left",
          //   label: "教程",
          // },

          //
          // todo x: 使用自定义导航栏, 参见 sidebars.js
          //

          // todo x: 教程
          {
            type: "docSidebar",
            position: "left",
            sidebarId: "guide",
            label: "文档",
          },
          {
            type: "docSidebar",
            position: "left",
            sidebarId: "biz",
            label: "产品列表",
          },
          {
            type: "docSidebar",
            position: "left",
            sidebarId: "service",
            label: "微服务列表",
          },


          // todo x: 生态
          {
            type: "docSidebar",
            position: "left",
            sidebarId: "ecosystem",
            label: "生态",
          },


          // TODO X: 博客
          {to: "/blog", label: "博客", position: "left"},

          {
            href: "https://github.com/zeromicro/zeromall",
            label: "GitHub",
            position: "right",
          },

          // todo x: 多语言
          {
            type: "localeDropdown",
            position: "right",
            dropdownItemsAfter: [
              {
                type: "html",
                value: '<hr style="margin: 0.3rem 0;">',
              },
              {
                href: "https://github.com/zeromicro/zeromall",
                label: "Help Us Translate",
              },
            ],
          },
        ],
      },
      footer: {
        style: "dark",
        links: [
          {
            title: "Docs",
            items: [
              {
                label: "Tutorial",
                to: "docs/guide/intro",
              },
              {
                label: "Ecosystem",
                to: "docs/ecosystem/index",
              },
            ],
          },
          {
            title: "Community",
            items: [
              {
                label: "Stack Overflow",
                href: "https://stackoverflow.com/questions/tagged/docusaurus",
              },
              {
                label: "Discord",
                href: "https://discord.com/invite/MnDA9pfWAW",
              },
              // {
              //     label: "Twitter",
              //     href: "https://twitter.com/gossip_coder", // todo x:
              // },
            ],
          },
          {
            title: "More",
            items: [
              {
                label: "Blog",
                to: "/blog",
              },
              {
                label: "GitHub",
                href: "https://github.com/zeromicro/zeromall", // todo x
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} ZeroMall. Built with Docusaurus.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
        additionalLanguages: ['powershell', 'go', 'rust', 'c', 'python'],
      },
    }),
};

module.exports = config;
