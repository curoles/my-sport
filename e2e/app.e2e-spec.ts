import { MySportPage } from './app.po';

describe('my-sport App', () => {
  let page: MySportPage;

  beforeEach(() => {
    page = new MySportPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
