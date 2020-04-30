it('Click home button', () => {
  cy.visit('/')

  cy.get('form').submit()

  cy.get('.results p')
    .should('contain', 'Thanks for clicking the button!')
})

it('Click back button', () => {
  cy.visit('/results')

  cy.get('form').submit()

  cy.get('.container p')
    .should('contain', 'Docker and Cypress Tutorial')
})