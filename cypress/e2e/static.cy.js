describe('static files', () => {
    it('serves tailwind', () => {
        cy.request('/static/css/tailwind.css').then(
            (resp) => {
                expect(resp.headers["content-type"]).to.contain("text/css")
                expect(resp.body).to.match(/\/\*! tailwindcss v[0-9]*.[0-9]*.[0-9]* | MIT License | https:\/\/tailwindcss.com \*\//)
            }
        )
    })
})