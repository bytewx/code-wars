function domainName(url) {
  url = url.replace(/https?:\/\/(www\.)?/, '').replace(/www\./, '');
  
  url = url.split('.')[0];
  
  return url;
}
