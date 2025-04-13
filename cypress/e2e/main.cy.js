describe('basic site serving', () => {
  it('serves html page', () => {
    cy.visit('/')

    cy.title().should("eq", "Tanuki-Ledger")
  })
})